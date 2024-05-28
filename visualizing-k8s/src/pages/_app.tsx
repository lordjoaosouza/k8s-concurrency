import '@/styles/globals.css'

import { QueryClientProvider } from '@tanstack/react-query'
import type { AppProps } from 'next/app'

import { ThemeProvider } from '@/components/themes/theme-provider'
import { queryClient } from '@/services/react-query'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ThemeProvider attribute='class' defaultTheme='system' enableSystem disableTransitionOnChange>
      <QueryClientProvider client={queryClient}>
        <Component {...pageProps} />
      </QueryClientProvider>
    </ThemeProvider>
  )
}
