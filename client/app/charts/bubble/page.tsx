import  FormContainer  from "@/components/FormContainer";
import { TimescaleGraphForm } from "@/components/TimescaleForm"
export default async function Page(){

    return(
        <>
        <FormContainer child={TimescaleGraphForm} />
        </>
    )
}