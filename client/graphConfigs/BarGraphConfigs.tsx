export const DefaultBarOptions = {
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
}

let titles = ["Animal Types", "Countries", "Music", "Beverages", "Cars", "Hobbies", "Languages", "Flowers", "Colours"]
function FakeBarTitles(n: number) {
    let r = []
    for (let i = 0; i < n; i++){
        r[i] = titles[Math.floor(Math.random() * titles.length)]
    }
    return r
}
function FakeDataNumbers(n: number){
    let r = []
    for (let i = 0; i < n ; i++){
        r[i] = Math.floor(Math.random()*100)
    }
    return r
}
export function FakeGraphLabelsAndData(l: number){
    let labels = FakeBarTitles(5)
    let datasets = [{
        label: "Fake Data Label",
        data: FakeDataNumbers(5),
        backgroundColor: 'rgb(75, 192, 192)'
    }]
    return {labels, datasets}
}