import { useQuery } from '@tanstack/react-query'
import { LoaderCircle, Repeat } from 'lucide-react'
import Link from 'next/link'
import { useEffect, useState } from 'react'

import { executeRequests } from '@/api/requests'
import RequestsCard from '@/components/requests/requests-card'
import { Button } from '@/components/ui/button'
import { queryClient } from '@/services/react-query'

import { Label } from '../ui/label'

interface RequestsContainerProps {
  requests: number
  setIsRunning: (isRunning: boolean) => void
}

export default function RequestsContainer({ requests, setIsRunning }: RequestsContainerProps) {
  const [isContentPresent, setIsContentPresent] = useState<boolean>(false)

  const { data: requestsData } = useQuery({
    queryKey: ['requests-data'],
    queryFn: () => executeRequests({ requests }),
    refetchOnWindowFocus: false,
    enabled: isContentPresent,
  })

  useEffect(() => {
    setIsContentPresent(true)
  }, [isContentPresent])

  useEffect(() => {
    console.log(requestsData)
  }, [requestsData])

  return (
    <div className={'flex h-[440px] w-[784px] flex-col items-center gap-4'}>
      <div className={'flex w-full items-center justify-between'}>
        <h1>Stress testing with {requests} requests</h1>
        <Link href={'/'}>
          <Button
            disabled={!requestsData}
            className={'gap-1'}
            onClick={() => {
              setIsRunning(false)
              setIsContentPresent(false)
              queryClient.removeQueries({ queryKey: ['requests-data'], exact: true })
            }}
          >
            New test <Repeat className={'size-4'} />
          </Button>
        </Link>
      </div>

      <div className={'flex grow items-center justify-center gap-4'}>
        {!requestsData && (
          <>
            <Label>Requesting...</Label>
            <LoaderCircle className={'animate-spin'} />
          </>
        )}

        {requestsData && (
          <>
            <RequestsCard
              title={'Single Docker container'}
              description={'With no concurrency'}
              icon={'docker'}
              requestAmount={requests}
              timeInMilliseconds={requestsData.single.data.timeInMilliseconds}
              hits={requestsData.single.data.hits}
              fails={requestsData.single.data.fails}
            />

            <RequestsCard
              title={'Kubernetes cluster'}
              description={'With concurrency'}
              icon={'k8s'}
              requestAmount={requests}
              timeInMilliseconds={requestsData.cluster.data.timeInMilliseconds}
              hits={requestsData.cluster.data.hits}
              fails={requestsData.cluster.data.fails}
            />
          </>
        )}
      </div>
    </div>
  )
}
