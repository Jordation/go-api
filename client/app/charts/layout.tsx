import Link from "next/link"

export default function GroupedBarLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
      <>
        {children}
        <Link href={"charts/groupedbar"}>to grouped bar </Link>
        <Link href={"charts/"}>to blah</Link>
      </> 
  )
}
