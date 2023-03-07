"use client"
import {use, useEffect, useState} from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
ChartJS.register(CategoryScale,LinearScale,BarElement,Title,Tooltip,Legend);
import { Bar } from 'react-chartjs-2';
import { useSearchParams } from 'next/navigation';


export const GroupedBarGraph = (props) => {

console.log(props.data)

    return (
            <div className='graphArea'>
                graph area
                {props.data}
            </div>
    )
}