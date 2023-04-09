"use client"
import {useForm} from "react-hook-form"
import { useRouter } from 'next/navigation';
import SelectFormItem from "./MultiDropdown";
import { FormProps } from "./FormContainer";
import { GroupedBarGraph } from "./GroupedBarGraph";
import { useState } from "react";

export const GroupedBarForm = ({props}: {props: FormProps}) => {
  
    const {register, handleSubmit, reset} = useForm()
    const [ChartData, setChartData] = useState(null)
    
    const onSubmit = (data:any) => {
        console.log("Sending request with params: ", data)
        const res = fetch("http://localhost:8000/GetGroupedBar", {
            method: "POST",
            body: JSON.stringify(data)
        })
        .then(res => res.json())
        .then(data => {
            setChartData(data)
            console.log(data)
        })
        .catch(err => console.log(err))
    }

    const resetButtonClick = () => reset()

    const RequestBar1 = () => {
        const res = fetch("http://localhost:8000/GetGroupedBar/1", {method: "GET"})
        .then (res => res.json())
        .then(data => setChartData(data))
        .catch(err => console.log(err))
    }
    const RequestBar2 = () => {
        const res = fetch("http://localhost:8000/GetGroupedBar/2", {method: "GET"})
        .then (res => res.json())
        .then(data => setChartData(data))
        .catch(err => console.log(err))
    }
    const RequestBar3 = () => {
        const res = fetch("http://localhost:8000/GetGroupedBar/3", {method: "GET"})
        .then (res => res.json())
        .then(data => setChartData(data))
        .catch(err => console.log(err))
    }
    const RequestBar4 = () => {
        const res = fetch("http://localhost:8000/GetGroupedBar/4", {method: "GET"})
        .then (res => res.json())
        .then(data => setChartData(data))
        .catch(err => console.log(err))
    }
    const RequestBar5 = () => {
        const res = fetch("http://localhost:8000/GetGroupedBar/5", {method: "GET"})
        .then (res => res.json())
        .then(data => setChartData(data))
        .catch(err => console.log(err))
    }

    return(
        <>
        <div className="formArea">
            form zoneature
            <form id="form" onSubmit={handleSubmit(onSubmit)}>
                <SelectFormItem register={register} field={"IS_filters.players"} list={props.players} multiple={true}/>
                <SelectFormItem register={register} field={"IS_filters.teams"} list={props.teams} multiple={true}/>
                <SelectFormItem register={register} field={"IS_filters.agents"} list={props.agents} multiple={true}/>
                <SelectFormItem register={register} field={"IS_filters.maps"} list={props.maps} multiple={true}/>
                <input type="text" {...register("side")}/>
                <br/>
                <SelectFormItem register={register} field={"x_target"} list={props.x_values} multiple={false}/>
                <SelectFormItem register={register} field={"x_groups_target"} list={props.x_values} multiple={false}/>
                <SelectFormItem register={register} field={"y_target"} list={props.y_values} multiple={false}/>
                <br/>
                <input type="number" {...register("min_dataset_size")}/>
                <input type="number" {...register("max_dataset_amount")}/>
                average results?
                <input type="checkbox" {...register("average_results")} value="true"/>
            </form>
            <button form="form">submit</button>
            
            <button onClick={resetButtonClick}>reset</button>

            <button onClick={RequestBar1}>1</button>
            <button onClick={RequestBar2}>2</button>
            <button onClick={RequestBar3}>3</button>
            <button onClick={RequestBar4}>4</button>
            <button onClick={RequestBar5}>5</button>

        </div>
        {ChartData && <GroupedBarGraph datasets={ChartData?.datasets} labels={ChartData?.labels}/>}
        </>
    )
}