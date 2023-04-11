package stats

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	GetResults = `select comp, 
	SUM(CASE WHEN won = 1 THEN 1 else 0 END) as wins, 
	SUM(CASE WHEN won = 0 THEN 1 else 0 END) as losses 
	FROM agent_comps
`
)

func (c *CompStats) Total() int64 { return c.Ws + c.Ls }

type CompStats struct {
	Comp string
	Ws   int64
	Ls   int64
}

func DerefIntegerInterface(v any) int64 {
	freqPtr := v.(*interface{})
	return (*freqPtr).(int64)
}

func GetCompsQuery(req *ListCompsRequest) (stmt string, args []interface{}) {
	var clauses []string
	stmt = GetResults

	if req.MapIds != nil {
		clauses = append(clauses, "map_id IN (?)")
		args = append(args, req.MapIds)
	}

	if req.INfilters.Agents != nil {
		for _, v := range req.INfilters.Agents {
			clauses = append(clauses, "comp LIKE (?)")
			args = append(args, "%"+v+"%")
		}
	}

	if req.NOTINfilters.Agents != nil {
		for _, v := range req.NOTINfilters.Agents {
			clauses = append(clauses, "comp NOT LIKE (?)")
			args = append(args, "%"+v+"%")
		}
	}

	if len(clauses) != 0 {
		stmt += " WHERE " + strings.Join(clauses, " AND ")
	}
	if req.Hc != nil {
		stmt += AddHavingClause(req.Hc)
	}
	log.Info("[GET COMPS Q, ARGS]: ", stmt, args)
	return stmt, args
}

func GetCompsWinratePickrate(dbConn *gorm.DB, req *ListCompsRequest) ([]CompStats, error) {
	var (
		dest []map[string]interface{}
		res  = make([]CompStats, 0)
	)
	req.Hc = &HavingClause{
		Targets: []string{"comp"},
		Count:   req.MinDatasetSize,
	}
	stmt, args := GetCompsQuery(req)
	log.Info("[GETTING WR+PR]")
	if err := dbConn.Table("agent_comps").Raw(stmt, args...).Scan(&dest).Error; err != nil {
		log.Error("[GET WR+PR]: error ", err)
	}

	for i := range dest {
		res = append(res, CompStats{
			Comp: dest[i]["comp"].(string),
			Ws:   DerefIntegerInterface(dest[i]["wins"]),
			Ls:   DerefIntegerInterface(dest[i]["losses"]),
		})
	}

	return res, nil
}
