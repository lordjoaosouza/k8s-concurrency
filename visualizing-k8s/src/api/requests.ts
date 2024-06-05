import { api } from '@/services/api'

interface RequestProps {
  requests: number
}

interface RequestData {
  single: {
    data: {
      timeInMilliseconds: number
      hits: number
      fails: number
    }
  }

  cluster: {
    data: {
      timeInMilliseconds: number
      hits: number
      fails: number
    }
  }
}

export async function executeRequests({ requests }: RequestProps) {
  const serverIp = process.env.NEXT_PUBLIC_SERVER_IP

  const singleResponse = await api.post('/picalc/testapi', {
    ip: `http://${serverIp}:8000/picalc/1000000`,
    quantity: requests,
  })

  const clusterResponse = await api.post('/picalc/testapi', {
    ip: `http://${serverIp}:8080/picalc/1000000`,
    quantity: requests,
  })

  return {
    single: { ...singleResponse.data },
    cluster: { ...clusterResponse.data },
  } as RequestData
}
