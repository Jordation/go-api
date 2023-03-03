
import "./global.css"


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
        <div className="header">header</div>
        <div className="mainContent">{children}</div>
        <div className="rightBar">right side bar</div>
    
      </body>
    </html>
  )
}
