"use client"
import {useForm} from "react-hook-form"
import { useRouter } from 'next/navigation';
import MultiDropdown from "./MultiDropdown";
import { ListUniqueStats } from "@/utils/getStats";

export default function GroupedBarForm({ formData }:object){
  
    const {register, handleSubmit} = useForm()
    const router = useRouter();

    // converts form output to query params
    const onSubmit = (data:any) => {
        console.log("data ", data)
        let url = new URL("http://localhost:3000/charts/groupedbar");
        let params = new URLSearchParams();
        for (const [key, value] of Object.entries(data)) {
            console.log(typeof value)
            if (!Array.isArray(value)) {
                for (const [key2, value2] of Object.entries(value)) {
                    params.append(key+"."+key2, value2)
                }
            } else {
                params.append(key, value)
            }
        url.search = params.toString();
        router.push(url.toString())
    }
    }
    
    const resetButtonClick = (data:any) => router.push("http://localhost:3000/charts/groupedbar")

    return(
        <>
        <div className="formArea">
            form zoneature
            <form id="form" onSubmit={handleSubmit(onSubmit)}>
                <MultiDropdown register={register} field={"filters.players"} list={formData.players}/>
                <MultiDropdown register={register} field={"filters.teams"} list={formData.teams}/>
                <MultiDropdown register={register} field={"filters.agents"} list={formData.agents}/>
                <MultiDropdown register={register} field={"filters.mapnames"} list={formData.maps}/>
                <br/>
                <MultiDropdown register={register} field={"filters_NOT.players"} list={formData.players}/>
                <MultiDropdown register={register} field={"filters_NOT.teams"} list={formData.teams}/>
                <MultiDropdown register={register} field={"filters_NOT.agents"} list={formData.agents}/>
                <MultiDropdown register={register} field={"filters_NOT.mapnames"} list={formData.maps}/>
            </form>
            <button form="form">submit</button>
            <button onClick={resetButtonClick}>reset</button>
        </div>

        </>
    )
}