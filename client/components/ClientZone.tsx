import GraphWithParams from "./GraphHolder"
import GraphForm from "./Form"

let query = {"agents": "",
"mapnames": "Icebox, Bind, Ascent, Haven, Breeze",
"players": "",
"teams": "",
"side": "C",
"average_rows_to_groups": "yer",
"order_by_y_target": "yes",
"min_dataset_size": "2",
"y_target": "kills",
"x_target": "map_name",
"x2_target": "agent",
"max_dataset_width": "5",
"query_level": "2"
}

export default function ClientZone({ params }: any){
    let query = params.query
    return(
        <>
        <div>
            {"params are: " + params.query}
            <GraphForm/>
            <GraphWithParams query={query}/>
        </div>
        </>
    )
}