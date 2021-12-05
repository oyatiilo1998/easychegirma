import FullPageLoader from "../../components/FullPageLoader"
import Sidebar from "../../components/Sidebar/DashboardSidebar"
import "./style.scss"

const MainLayout = ({children}) => {

  return (
    <div className="MainLayout">
      <div className="sidebar-side">
        <Sidebar />
      </div>
      <div className="content-side">
        <FullPageLoader />
        {children}
      </div>
    </div>
  )
}

export default MainLayout
