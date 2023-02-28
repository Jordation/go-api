"use client"
import {use} from 'react';
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

const options = {
    plugins: {
        title: {
            display: true,
            text: 'Chart.js Bar Chart - Stacked',
        },
    },
    responsive: true,
    interaction: {
        mode: 'index' as const,
        intersect: false,
    },
};
async function DealWithParams(params: any){
    await new Promise(resolve => setTimeout(resolve, 1000));
    return params.query.split("")
}

export default function GraphWithParams({ params }: any) {
    ChartJS.register(
        CategoryScale,
        LinearScale,
        BarElement,
        Title,
        Tooltip,
        Legend
    );
    const labels = ["aye", "two", "four"]
    const chartData = use(DealWithParams(params))
    const data = {
        labels,
        datasets: [
            {
                label: params.query,
                data: chartData,
                backgroundColor: 'rgb(75, 192, 192)'
            }
        ]
    }
    return (
        <>
            <div>i got these params: {params.query}</div>
            <Bar options={options} data={data} />
        </>

    )
}
