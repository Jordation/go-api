import GroupedBarForm from "@/components/GroupedBarForm"

export default function GroupedBarLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
      <>
        <GroupedBarForm />
        {children}
      </>
  )
}
