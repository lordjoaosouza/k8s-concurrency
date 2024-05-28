import axios from 'axios'

export const single = axios.create({
  baseURL: 'http://localhost:8080',
})

export const cluster = axios.create({
  baseURL: 'http://localhost:8080',
})
