import { Paragraph } from "../elements/Typography"
import { Brand } from "../fragments/Brand"
import { FeedbackForm } from "../fragments/FeedbackForm"

export const FooterLayout = (props) => {
    return (
        <footer className={props.className}>
            <div className="footer-container">
                <div className="footer-left">
                    <div className="footer-left-1">
                        <Brand className="footer-title"></Brand>
                        <Paragraph id="slogan">Wear Confidence, Wear Cays</Paragraph>
                    </div>
                        <Paragraph id="license">© Copyright 2024 CaysFashion</Paragraph>
                </div>
                <div className="footer-right">
                    <FeedbackForm/>
                </div>
            </div>
        </footer>
    )
}