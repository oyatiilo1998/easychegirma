import "./style.scss"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import {
  faPhoneAlt,
  faShareAlt,
  faMapMarkerAlt,
} from "@fortawesome/free-solid-svg-icons"
import {
  faInstagram,
  faTelegram,
  faFacebook,
} from "@fortawesome/free-brands-svg-icons"

const Footer = () => {
  return (
    <div className="Footer">
      <div className="footer-block container">
        <div className="column">
          <h3 className="title">О НАС</h3>
          <div className="row">Новости</div>
          <div className="row">Гарантия</div>
          <div className="row">Доп. гарантия</div>
          <div className="row">Акции</div>
        </div>
        <div className="column">
          <h3 className="title">Информация</h3>
          <div className="row">Новости</div>
          <div className="row">Бонусная система</div>
          <div className="row">Контакты</div>
          <div className="row">Акции</div>
        </div>
        <div className="column"></div>

        <div className="column">
          <div className="contact-row">
            <div className="icon">
              <FontAwesomeIcon icon={faPhoneAlt} />
            </div>
            <div className="info-side">
              <div className="subtitle">Возник вопрос? Звоните</div>
              <div className="info">+998 99 123-45-67</div>
            </div>
          </div>

          <div className="contact-row">
            <div className="icon">
              <FontAwesomeIcon icon={faMapMarkerAlt} />
            </div>
            <div className="info-side">
              <div className="subtitle">Наш адрес</div>
              <div className="info">Г. Ташкент</div>
            </div>
          </div>

          <div className="contact-row">
            <div className="icon">
              <FontAwesomeIcon icon={faShareAlt} />
            </div>
            <div className="info-side">
              <div className="subtitle">Мы в соцсетях</div>
              <div className="info">Подписывайтесь</div>
              <div className="extra">
                <div className="social-icon"><FontAwesomeIcon icon={faInstagram} /></div>
                <div className="social-icon"><FontAwesomeIcon icon={faTelegram} /></div>
                <div className="social-icon"><FontAwesomeIcon icon={faFacebook} /></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Footer
