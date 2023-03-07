import  FormContainer  from "@/components/FormContainer";
import { GroupedBarForm } from "@/components/GroupedBarForm";
import { GroupedBarGraph } from "@/components/GroupedBarGraph";

export default async function Page(){

    return(
        <>
        <FormContainer child={GroupedBarForm}/>
        <GroupedBarGraph/>
        </>
    )
}