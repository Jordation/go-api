import FormContainer from "@/components/FormContainer";
import GroupedBarForm from "@/components/GroupedBarForm";
import GroupedBarGraph from "@/components/GroupedBarGraph";
async function getUniqueStatLists(target: string){
    const res = await fetch(`http://localhost:8000/ListUniqueStats/${target}`, {
        method: "GET",})
    return res.json()
}
export default async function Page(){

    const playerList = await getUniqueStatLists("player_name")
    const teamList = await getUniqueStatLists("team")
    const agentList = await getUniqueStatLists("agent")
    const mapList = await getUniqueStatLists("map_name")

    const [players, teams, agents, maps] = await Promise.all([playerList.data, teamList.data, agentList.data, mapList.data]) 
    const lists = {players, teams, agents, maps}
    return(
        <>
        
        <GroupedBarForm formData={lists}/>
        <GroupedBarGraph/>
        </>
    )
}