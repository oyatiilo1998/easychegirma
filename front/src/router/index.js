import { Redirect, Route, Switch, useLocation } from "react-router"
import { useTransition, animated } from "react-spring"
import MainLayout from "../layouts/MainLayout"
import { routesList } from "./routesList"

const AppRouter = () => {
  const location = useLocation()
  const transitions = useTransition(location, {
    from: { opacity: 0 },
    enter: { opacity: 1 },
    leave: { opacity: 0 },
  })

  return (
    <Switch>
      <MainLayout>
        {transitions((props, item) => ( 
          <animated.div style={props}>
            <div style={{ position: "absolute", width: "100%" }}>
              <Switch location={item}>
                {
                  routesList.map(route => (
                    <Route
                      {...route}
                    />
                  ))
                }
                <Redirect to="/" />
              </Switch>
            </div>
          </animated.div>
        ))}
      </MainLayout>
    </Switch>
  )
}

export default AppRouter
