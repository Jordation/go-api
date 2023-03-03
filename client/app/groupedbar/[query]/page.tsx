import GroupedBarGraph from "@/components/GroupedBarGraph";
import { DefaultBarOptions, FakeGraphLabelsAndData } from '@/graphConfigs/BarGraphConfigs';

export default function GraphWithParams({ params }: any) {

    let options = DefaultBarOptions
    let data = FakeGraphLabelsAndData(5)
    return (
        <>
            <GroupedBarGraph data={data} options={options} />
        </>

    )
}
