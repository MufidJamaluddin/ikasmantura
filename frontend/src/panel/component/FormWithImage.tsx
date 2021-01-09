import React, {PureComponent} from "react";

export abstract class FormWithImage extends PureComponent<any, { image:any }> {

    protected constructor(props) {
        super(props);
        this.state = {
            image: null
        }
        this.transform = this.transform.bind(this)
        this.dropImage = this.dropImage.bind(this)
    }

    abstract transformData(image, data): any

    transform(data): any {
        let image = this.state.image
        return this.transformData(image, data)
    }

    dropImage(file) {
        console.log('image:', file);
        this.setState({ image: file })
    }

}
