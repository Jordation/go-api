
import "./global.css"
import Link from 'next/link';

export const metadata = {
  title: 'graphs mate',
  description: 'its my new site g',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="main">
        
        <div className="leftBar">left side bar</div>
        <div className="header">
          
          <Link href={"/charts/groupedbar"}> grouped bar </Link>
          <Link href={"/charts"}> charts </Link>
          <Link href={"/stats"}> stats </Link>
          <Link href={"/"}> home </Link>
          header
        </div>
        <div className="mainContent">{children}</div>
        <div className="rightBar">right side bar</div>
    
      </body>
    </html>
  )
}
