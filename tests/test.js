import order from "k6/x/order"

export const options = {
    vus: 1,
    iterations: 1,
}

export default function () {
    var body = "{\"orderId\":1619286267885357876}"
    console.log(body)

    var orderId = order.getOrderID(body)
    console.log(orderId)
}