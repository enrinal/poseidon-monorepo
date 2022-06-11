import axios from "axios";
import NodeCache from "node-cache";

const cache = new NodeCache({ stdTTL: 200 });

export default class CurrencyConverterRepository {
  public static async fetchCurrencyConverter() {
    if (cache.has("currency")) {
      return cache.get("currency");
    }
    try {
      let config = {
        headers: {
          apikey: process.env.CURRENCY_CONVERTER_API_KEY as string,
        },
      };
      const d = new Date();
      const dateString =
        d.getFullYear() +
        "-" +
        ("0" + (d.getMonth() + 1)).slice(-2) +
        "-" +
        ("0" + d.getDate()).slice(-2);

      const url = (process.env.CURRENCY_CONVERTER_URL as string)
        .replace("{start_date}", dateString)
        .replace("{end_date}", dateString);

      const response = await axios.get(url, config);
      cache.set(
        "currency",
        response.data.quotes.USDIDR.start_rate,
        60 * 60 * 24
      );
      return response.data.quotes.USDIDR.start_rate;
    } catch (error) {
      console.log(error);
      throw error;
    }
  }
}
