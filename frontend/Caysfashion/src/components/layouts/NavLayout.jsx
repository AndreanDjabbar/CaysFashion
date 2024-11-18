import { NavTitle } from "../fragments/NavTitle"

export const NavLayout = (props) => {
    return (
        <nav className={props.className}>
            <NavTitle/>
            <div className="nav-link">
                {props.children}
            </div>
        </nav>
    )
}