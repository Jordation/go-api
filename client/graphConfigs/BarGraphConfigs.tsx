export const DefaultBarOptions = {
    plugins: {
        tooltip: {
            filter: function(tooltipItem) {
                if(tooltipItem.parsed.y == null){
                    return false
                } else {
                    return true
                }
            }
        },
        title: {
            display: true,
            text: 'Chart.js Bar Chart - Stacked',
        },
    },
    interaction: {
        mode: 'index' as const,
        intersect: false,
    },
    responsive: true,
    skipNull: true
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