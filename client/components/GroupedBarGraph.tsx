"use client";
import { use, useEffect, useState, useRef } from "react";
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

import { Bar, getElementsAtEvent } from "react-chartjs-2";
import { useSearchParams } from "next/navigation";
import { DefaultBarOptions } from "@/graphConfigs/BarGraphConfigs";
	
export const GroupedBarGraph = (props) => {
	const chartRef = useRef(null);
	useEffect(() => {
		const chart = chartRef.current
		console.log(chartRef)
		chart.update()
	}, [props]);
	return (
		<div className="graphArea">
			graph area
			<Bar ref={chartRef} datasetIdKey="id" data={props} options={DefaultBarOptions} />
			<button onClick={() => console.log(props)}>click</button>
		</div>
	);
};
	