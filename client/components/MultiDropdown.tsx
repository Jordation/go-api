import { useEffect, useState } from "react"

export default function SelectFormItem({register, field, list, multiple}: any) {

    return (
        <>
            <select id={field} {...register(field)} multiple={multiple}>
                {list.sort().map((opt, i) => {
                    return <option key={i} value={opt}>{opt}</option>
                })}
            </select>
        </>
    )
}