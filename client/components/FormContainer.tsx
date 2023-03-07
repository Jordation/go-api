import { ReactNode } from "react"

async function getUniqueStatLists(target: string){
    const res = await fetch(`http://localhost:8000/ListUniqueStats/${target}`, {
        method: "GET",})
        console.log("i did a fetch for: ", target)
    return res.json()
}

export interface FormProps {
        agents: string[],
        maps: string[],
        players: string[],
        teams: string[],
        x_values: string[],
        y_values: string[],
}
interface FormContainerProps {
    child: ReactNode
}

export default async function FormContainer(props: FormContainerProps){

    const playerList = await getUniqueStatLists("player_name")
    const teamList = await getUniqueStatLists("team")
    const agentList = await getUniqueStatLists("agent")
    const mapList = await getUniqueStatLists("map_name")
    const x_values = ["player_name", "team", "agent", "map_name"]
    const y_values = ["kills", "deaths", "assists", "rating", "kda", "adr", "kast", "hsp"]

    const [players, teams, agents, maps] = await Promise.all([playerList.data, teamList.data, agentList.data, mapList.data]) 
    
    const lists = {players, teams, agents, maps, x_values, y_values}
    
    let Form = props.child

    return(
        <>
        form container
        <div className="formArea">
            <Form props={lists}/>
        </div>
        </>
    )
}
