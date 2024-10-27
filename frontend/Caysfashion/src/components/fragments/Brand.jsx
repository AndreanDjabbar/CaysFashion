import { Logo } from "../elements/Logo"
import { Paragraph } from "../elements/Typography"

export const Brand = (props) => {
    return (
        <div className={props.className}>
            <Logo/>
            <Paragraph>CaysFashion</Paragraph>
        </div>
    )
}