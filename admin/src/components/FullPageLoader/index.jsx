import { CircularProgress } from "@mui/material"
import { useSelector } from "react-redux"
import "./style.scss"

const FullPageLoader = ({isActive}) => {
  const isVisible = useSelector(state => state.loader.fullPageLoader)
  
  if(!isVisible && !isActive) return null

  return (
    <div className="FullPageLoader">
      <CircularProgress color="primary" />
    </div>
  )
}

export default FullPageLoader
