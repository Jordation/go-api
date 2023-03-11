"use client"

import { useState } from "react"
import { useForm } from "react-hook-form"
import SelectFormItem from "./MultiDropdown"
import RowResults from "./RowResults"

export default function StatsViewForm ({props}){
    const [Rows, setRows] = useState(null)
    const {register, handleSubmit, reset} = useForm()
    const handleClick = (data) => {
        let res = fetch("http://localhost:8000/ListStats", {
            method: "POST",
            body: JSON.stringify(data)
        })
        .then(res => res.json())
        .then(res => {setRows(res.Stats)})
    }
    
    return(
        <>
        <form id="form" onSubmit={handleSubmit(handleClick)}>
                <SelectFormItem register={register} field={"filters.players"} list={props.players} multiple={true}/>
                <SelectFormItem register={register} field={"filters.teams"} list={props.teams} multiple={true}/>
                <SelectFormItem register={register} field={"filters.agents"} list={props.agents} multiple={true}/>
                <SelectFormItem register={register} field={"filters.mapnames"} list={props.maps} multiple={true}/>
                <input type="text" {...register("filters.side")}/>
                <br/>
            </form>
            <button form="form">submit</button>
        {Rows && <RowResults Rows={Rows}/>}
        </>
    )
}