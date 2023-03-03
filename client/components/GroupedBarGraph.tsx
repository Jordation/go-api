"use client"
import {use, useState} from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';
ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
    );
    


export default function GroupedBarGraph( {data, options}: any ) {

    //fetch("http://localhost:8000/testing", {method: "POST", body: JSON.stringify({})})
    return (
            <div className='graphArea'>
                <Bar options={options} data={data} />
            </div>
    )
}
