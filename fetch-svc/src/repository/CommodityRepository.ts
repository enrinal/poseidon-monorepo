import axios from "axios";

export default class CommodityRepository {
  public static async fetchCommodity() {
    try {
      const response = await axios.get(
        process.env.EFISHERY_STORAGES_URL as string
      );
      return response.data;
    } catch (error) {
      console.log(error);
      throw error;
    }
  }
}
