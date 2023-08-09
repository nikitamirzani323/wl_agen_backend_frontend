export default function (node, p) {
    let params = p;
    console.log("Attaching click ", node, params);

    node.addEventListener("click", (e) => {
        console.log("You clicked ", params, node, e)
    });
    return {
        update(p) {
            params = p
            console.log("Update action",params)
        },
        destroy() {
            console.log("Detaching / destroy click")
        }
    }
}