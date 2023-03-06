import { useEffect, useState } from "react"

export default function MultiDropdown({register, field, list}: any) {

    return (
        <>
            <select {...register(field)} multiple={true}>
                {list.map((opt, i) => {
                    return <option key={i} value={opt}>{opt}</option>
                })}
            </select>
        </>
    )
}