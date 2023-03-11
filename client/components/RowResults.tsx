
interface RowResultsProps {
    Rows: object[]
}
interface dbRow {
    PlayerName: string
    Team: string
    MapName: string
    Agent: string
    Rating: number
    ACS: number
    Kills: number
    Deaths: number
    Assists: number
    KAST: number
    ADR: number
    HSP: number
    FK: number
    FD: number
}


export default function RowResults(props: RowResultsProps){
    console.log("here are props", props.Rows)

    const rows = props.Rows.map((row) => {
        let vals = Object.values(row).map((value, i) => {
            return <div key={i+value}>{value}</div>
        })
        return <div className="rowResult" key={JSON.stringify(row)}>{vals}</div>})

    return(
        <>
        <div className="rowResults">
        {rows}
        </div>
        </>
    )
}