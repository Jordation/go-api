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


  const e = 2.71828
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
        let rad =context.dataset.data[context.dataIndex].r;
        return 10 + ((1 / (1 + e**(-0.1 * (rad - 50)))) * 70)
      },
      pointHoverRadius: function(context) {
        return context.dataset.data[context.dataIndex].r*0.75;
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
      suggestedMin: 40,
      suggestedMax: 70,
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
	