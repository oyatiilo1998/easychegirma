import axios from "axios";
import { store } from "../redux/store";
export const baseURL = "https://api.chegirma.uz"

const reqGenerator = axios.create({
  baseURL,
  timeout: 100000,
})

const errorHandler = (error, hooks) => {
  if (error?.response) {
    if(error.response?.data?.message) {
      // store.dispatch(showAlert(error.response.data.message))
    }

    if (error?.response?.status === 403) {

    }
    else if ( error?.response?.status === 401) {
      // store.dispatch(logout())
    }
  }

  return Promise.reject(error.response)
}

reqGenerator.interceptors.request.use(
  config => {
    const token = store.getState().auth.token
    if(token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },

  error => errorHandler(error)
)

reqGenerator.interceptors.response.use(response => response.data , errorHandler)

export default reqGenerator
