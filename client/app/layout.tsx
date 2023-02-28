
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
      <body>{children}</body>
    </html>
  )
}
