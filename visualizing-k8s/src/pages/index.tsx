import { BarChart2, Flame } from 'lucide-react'
import { Inter } from 'next/font/google'
import { useState } from 'react'

import RequestsContainer from '@/components/requests/requests-container'
import { ModeToggle } from '@/components/themes/mode-toggle'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const [isRunning, setIsRunning] = useState<boolean>(false)
  const [requests, setRequests] = useState<number>(100)

  return (
    <main className={`flex h-screen flex-col items-center justify-between p-24 ${inter.className}`}>
      <header className={'flex w-full items-center justify-between'}>
        <div className={'flex items-center gap-2'}>
          <BarChart2 /> Visualizing K8s
        </div>
        <ModeToggle />
      </header>

      <div className={'flex h-full w-full flex-col items-center justify-center gap-8'}>
        {!isRunning && (
          <div className={'flex flex-col items-center gap-4'}>
            <div className={'flex w-full items-center gap-2'}>
              <h1 className={'text-nowrap'}>Stress test with</h1>
              <Input
                placeholder={'...'}
                type={'number'}
                className={'h-8 w-20'}
                defaultValue={100}
                value={requests}
                onChange={(e) => setRequests(Number(e.target.value))}
              />
              <h1>requests</h1>
            </div>

            <Button
              className={'gap-1'}
              onClick={() => {
                setIsRunning(!isRunning)
              }}
            >
              Fire <Flame className={'size-4'} />
            </Button>
          </div>
        )}

        {isRunning && <RequestsContainer requests={requests} setIsRunning={setIsRunning} />}
      </div>
    </main>
  )
}
