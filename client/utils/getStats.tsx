import { cache } from "react";

export const ListUniqueStats = cache(async (target: string) => {
    const res = await fetch(`http://localhost:8000/ListUniqueStats/${target}`, {
        method: 'GET',
    })
    .then(res => res.json())
    .catch(err => console.log(err))
    return res
});