import ClientZone from "@/components/ClientZone";

export default function GraphWithParams({ params }: any) {
 
    return (
        <>
        <div>
            i am the graph page after a query
            <ClientZone params={params} />
        </div>
        </>

    )
}
