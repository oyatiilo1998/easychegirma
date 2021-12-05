import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import Logo from "../Logo"
import "./style.scss"
import {
  faSearch,
} from "@fortawesome/free-solid-svg-icons"

const Header = () => {
  return (
    <div className="Header">
      <div className="header-block container">
        <Logo />
        <div className="searchInput">
          <input  type="text" placeholder="Поиск" />
          <div className="search-btn"><FontAwesomeIcon icon={faSearch} /></div>
        </div>
      </div>
      <div className="navbar-wrapper">
        <div className="navbar container">
          <div className="navbar-item">Электроника</div>
          <div className="navbar-item">Одежда</div>
          <div className="navbar-item">Для дома</div>
          <div className="navbar-item">Еда и напитки</div>
          <div className="navbar-item">Телефоны</div>
        </div>
      </div>
    </div>
  )
}

export default Header
