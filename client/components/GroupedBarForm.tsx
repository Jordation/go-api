"use client"
import {useForm} from "react-hook-form"
import { useRouter } from 'next/navigation';
import SelectFormItem from "./MultiDropdown";
import { FormProps } from "./FormContainer";
import { GroupedBarGraph } from "./GroupedBarGraph";
import { useState } from "react";
export const GroupedBarForm = ({props}: {props: FormProps}) => {
  
    const {register, handleSubmit, reset} = useForm()
    const [ChartData, setChartData] = useState("")
    
    const onSubmit = (data:any) => {
        console.log("Sending request with params: ", data)
        const res = fetch("http://localhost:8000/graphs/groupedBar", {
            method: "POST",
            body: JSON.stringify(data)
        })
        .then(res => res.json())
        .then(data => {setChartData(data)})
        .catch(err => console.log(err))
    }

    const resetButtonClick = () => reset()

    return(
        <>
        <div className="formArea">
            form zoneature
            <form id="form" onSubmit={handleSubmit(onSubmit)}>
                <SelectFormItem register={register} field={"filters.players"} list={props.players} multiple={true}/>
                <SelectFormItem register={register} field={"filters.teams"} list={props.teams} multiple={true}/>
                <SelectFormItem register={register} field={"filters.agents"} list={props.agents} multiple={true}/>
                <SelectFormItem register={register} field={"filters.mapnames"} list={props.maps} multiple={true}/>
                <br/>
                <SelectFormItem register={register} field={"x_target"} list={props.x_values} multiple={false}/>
                <SelectFormItem register={register} field={"x_groups_target"} list={props.x_values} multiple={false}/>
                <SelectFormItem register={register} field={"y_target"} list={props.y_values} multiple={false}/>
                <br/>
                <input type="number" {...register("min_dataset_size")}/>
                <input type="number" {...register("max_dataset_amount")}/>
                <input type="text" {...register("average_results")}/>
            </form>
            <button form="form">submit</button>
            
            <button onClick={resetButtonClick}>reset</button>
        </div>
        <GroupedBarGraph data={ChartData} />
        </>
    )
}