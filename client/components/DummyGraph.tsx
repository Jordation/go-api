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

export const DummyGraph = () => {



  return (
    <div className="graphArea">

    </div>
  );
};
