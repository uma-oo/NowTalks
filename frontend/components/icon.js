import { createElement } from "/frontend/utils.js"

export function createIcon(name, type) {
    switch (type) {
        case "like":
            const svgNS = "http://www.w3.org/2000/svg";
            const svg = document.createElementNS(svgNS, "svg");

            svg.setAttribute("width", "24");
            svg.setAttribute("height", "24");
            svg.setAttribute("viewBox", "0 0 24 24");
            svg.setAttribute("fill", "none");
            svg.setAttribute("xmlns", svgNS);

            // Create the path element
            const path = document.createElementNS(svgNS, "path");

            path.setAttribute("d", "M12.62 20.81C12.28 20.93 11.72 20.93 11.38 20.81C8.48 19.82 2 15.69 2 8.68998C2 5.59998 4.49 3.09998 7.56 3.09998C9.38 3.09998 10.99 3.97998 12 5.33998C13.01 3.97998 14.63 3.09998 16.44 3.09998C19.51 3.09998 22 5.59998 22 8.68998C22 15.69 15.52 19.82 12.62 20.81Z");
            path.setAttribute("stroke", "#47307B");
            path.setAttribute("stroke-width", "1.5");
            path.setAttribute("stroke-linecap", "round");
            path.setAttribute("stroke-linejoin", "round");

            // Append the path to the SVG
            svg.appendChild(path);
            return svg

        default:
            let icon = createElement('img', 'icon')
            icon.src = `/frontend/assets/icons/${name}.svg`
            return icon

    }
}