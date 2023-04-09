"use client"
import {useForm} from "react-hook-form"
import { useRouter } from 'next/navigation';
import SelectFormItem from "./MultiDropdown";
import { FormProps } from "./FormContainer";
import { BubbleTest } from "./BubbleTest";
import { useState } from "react";
import { Chart } from "chart.js";

export const TimescaleGraphForm = ({props}: {props: FormProps}) => {
  
    const [ChartData, setChartData] = useState(null)

    const RequestBar1 = () => {
        const res = fetch("http://localhost:8000/GetTimescaleGraph/1", {method: "GET"})
        .then (res => res.json())
        .then(data => setChartData(data))
        .catch(err => console.log(err))
    }

    return(
        <>
        <div className="formArea">
            form zoneature

            <button onClick={RequestBar1}>1</button>

        </div>
        {ChartData && <BubbleTest data={ChartData}/>}
        </>
    )
}