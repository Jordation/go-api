
import FormContainer from "@/components/FormContainer"
import SelectFormItem from "@/components/MultiDropdown"
import StatsViewForm from "@/components/StatsViewForm"
import { useForm } from "react-hook-form"





export default function Page(){
    
    
    return (
        <>
        <div>stats viewer</div>
        <FormContainer child={StatsViewForm} />
        </>
    )
}