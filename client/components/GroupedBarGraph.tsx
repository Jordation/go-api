"use client";
import { use, useEffect, useState } from "react";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);
import { Bar } from "react-chartjs-2";
import { useSearchParams } from "next/navigation";
import { DefaultBarOptions } from "@/graphConfigs/BarGraphConfigs";

interface DataReturn {
        [key: string]: number[];
}

export interface GroupedBarProps {
    Labels: string[] | undefined;
    Data: DataReturn | undefined;
}

function rand_rgb() { // random colour
    let r = Math.floor(Math.random() * 256);
    let g = Math.floor(Math.random() * 256);
    let b = Math.floor(Math.random() * 256);
    return `rgb(${r}, ${g}, ${b})`;
}

function HandleData(props: GroupedBarProps){
  if (props.Data!=undefined){
    let labels = props.Labels
      let datasets = (Object.keys(props.Data)).map((key, i) => {
          let data = props.Data?.[key]
          for(let i=0 ; i<data.length;i++){
              if (data[i] == 0) data[i]=null;
          }
      return {
        id: i, 
        label: key, 
        data: props.Data?.[key], 
        backgroundColor: rand_rgb()
      }})
      return {labels: labels, datasets: datasets}
  }
}

export const GroupedBarGraph = (props: GroupedBarProps) => {
    const [ChartData, setChartData] = useState(null)
    const data = HandleData(props)
    useEffect(() => {
        setChartData(HandleData(props))
    }, [props])
    
  return (
    <div className="graphArea">
      graph area

      {ChartData && 
      <Bar datasetIdKey="id" data={ChartData} options={DefaultBarOptions} />}
    <button onClick={()=>console.log(props)}>click</button>
    </div>
  );
};
