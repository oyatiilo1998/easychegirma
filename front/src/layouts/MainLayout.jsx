import Footer from "../components/Footer"
import Header from "../components/Header"

const MainLayout = ({children}) => {
  return (
    <div className="MainLayout" >
      <Header />
      <div className="content-block" style={{ paddingTop: '120px' }} >
        {children}
      </div>
      
    </div>
  )
}

export default MainLayout
