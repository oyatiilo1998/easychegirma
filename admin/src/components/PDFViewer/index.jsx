import { CircularProgress } from "@mui/material"
import MGRPDFViewer from "mgr-pdf-viewer-react"

const PDFViewer = ({ url, isVisible }) => {
  if(!url) return null

  return (
    <div style={{  width: '100%', backgroundColor: '#838689', paddingTop: '30px' }} >
      <MGRPDFViewer
        document={{
          url,
        }}
        style={{width: '100%'}}
        loader={<CircularProgress />}
      />
    </div>
  )
}

export default PDFViewer
