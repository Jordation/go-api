"use client";
import { use, useEffect, useState, useRef } from "react";
import {
  Chart as ChartJS,
  LinearScale,
  PointElement,
  Tooltip,
  Legend,
} from 'chart.js/auto';
import { Line } from 'react-chartjs-2';
import 'chartjs-adapter-moment';

ChartJS.register(
  LinearScale,
  PointElement,
  Tooltip,
  Legend,
  );


function mapResults(inData) {
    const res = inData.datasets.map(result => ({
      label: result.label,
      data: result.data.map(item => ({
        x: new Date(item.x),
        y: item.y,
        r: item.r
      })),
      backgroundColor: result.backgroundColor,
      borderColor: result.borderColor,
      borderWidth: 1,
      pointRadius: function(context) {
        return context.dataset.data[context.dataIndex].r;
      },
      pointHoverRadius: function(context) {
        return context.dataset.data[context.dataIndex].r * 1.5;
      }
    }));
    console.log(res)
    return {datasets: res}
}

const options = {
  scales: {
    x: {
      type: 'time',
      time: {
        unit: 'month',
      },
    },
    y: {
      suggestedMin: 0,
      suggestedMax: 100,
    },
  },
};
import { Bar, getElementsAtEvent } from "react-chartjs-2";
	
export const BubbleTest = (props) => {
  const [data, setdata] = useState(null)
	useEffect(() => {
    if (props != null){
     setdata(mapResults(props.data))
    }
	}, [props]);

	return (
		<div className="graphArea">
			graph area
			{data && <Line data={data} options={options} />}
		</div>
	);
};
	