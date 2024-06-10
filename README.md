# Merge Hotel Data 

This project merges hotel data from multiple suppliers, cleans it, and delivers it via an API endpoint.

### Getting Started

Follow these instructions to set up and run the project on your local machine.

### Prerequisites

- Make sure you have [Docker Desktop](https://www.docker.com/products/docker-desktop) installed and running.

### Installation

1. **Clone the repository**:

            git clone https://github.com/avinara/merge-hotel-data.git
            cd merge-hotel-data
2. **Make sure Docker Desktop is running**:

    - Docker Desktop should be running to ensure the Docker daemon is switched on.

3. **Navigate to the project directory**:

            cd merge-hotel-data

4. **Build the Docker image**:

            docker-compose build

5. **Run the application**:
    
            docker-compose up

### Usage
   - Once the application is running, you can access the API at http://localhost:8080.

### API Endpoints
            localhost:8080/hotels?searchStr=destination_id&searchValue=1122
            localhost:8080/hotels?searchStr=id&searchValue=f8c9,iJhz

### Query Params

    `searchStr =  "id" , searchValue =  "f8c9,iJhz" `
    `searchStr =  "destination_id", searchValue = "1122"`
	
### Response Structure

The merged response from all the suppliers would look like below:
```
[
  {
    "id": "f8c9",
    "destination_id": 1122,
    "name": "Hilton Shinjuku Tokyo",
    "location": {
      "lat": 35.6926,
      "lng": 139.690965,
      "address": "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
      "city": "Tokyo",
      "country": "Japan"
    },
    "description": "This sleek high-rise property is 10 minutes' walk from Shinjuku train station, 6 minutes' walk from the Tokyo Metropolitan Government Building and 3 km from Yoyogi Park. The polished rooms offer Wi-Fi and flat-screen TVs, plus minibars, sitting areas, and tea and coffeemaking facilities. Suites add living rooms, and access to a club lounge serving breakfast and cocktails. A free shuttle to Shinjuku station is offered. There's a chic Chinese restaurant, a sushi bar, and a grill restaurant with an open kitchen, as well as an English pub and a hip cocktail lounge. Other amenities include a gym, rooftop tennis courts, and a spa with an indoor pool.",
    "amenities": {
      "general": [
        "pool",
        "business center",
        "dry cleaning",
        "bar",
        "wifi"
      ],
      "rooms": [
        "tv",
        "aircon",
        "minibar",
        "bathtub",
        "indoor pool"
      ]
    },
    "images": {
      "rooms": [
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i10_m.jpg",
          "description": "Suite"
        },
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i11_m.jpg",
          "description": "Suite - Living room"
        }
      ],
      "site": [
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i55_m.jpg",
          "description": "Bar"
        }
      ],
      "amenities": [
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i57_m.jpg",
          "description": "Bar"
        }
      ]
    },
    "booking_conditions": [
      "All children are welcome. One child under 6 years stays free of charge when using existing beds. There is no capacity for extra beds in the room.",
      "Pets are not allowed.",
      "Wired internet is available in the hotel rooms and charges are applicable. WiFi is available in the hotel rooms and charges are applicable.",
      "Private parking is possible on site (reservation is not needed) and costs JPY 1500 per day.",
      "When booking more than 9 rooms, different policies and additional supplements may apply.",
      "The hotel's free shuttle is offered from Bus Stop #21 in front of Keio Department Store at Shinjuku Station. It is available every 20-minutes from 08:20-21:40. The hotel's free shuttle is offered from the hotel to Shinjuku Train Station. It is available every 20-minutes from 08:12-21:52. For more details, please contact the hotel directly. At the Executive Lounge a smart casual dress code is strongly recommended. Attires mentioned below are strongly discouraged and may not permitted: - Night attire (slippers, Yukata robe, etc.) - Gym clothes/sportswear (Tank tops, shorts, etc.) - Beachwear (flip-flops, sandals, etc.) and visible tattoos. Please note that due to renovation works, the Executive Lounge will be closed from 03 January 2019 until late April 2019. During this period, guests may experience some noise or minor disturbances. Smoking preference is subject to availability and cannot be guaranteed."
    ]
  }
]
```

### Data Procurement Strategy

 For each supplier we have the following configs defined in the config.json
```
        {
            "name": "acme",
            "source": "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme",
            "response_format": {
                "id": "Id",
                "destination_id": "DestinationId",
                "name": "Name",
                "lat": "Latitude",
                "lng": "Longitude",
                "address": "Address",
                "city": "City",
                "country": "Country",
                "postal_code": "PostalCode",
                "description": "Description",
                "amenities": "Facilities"
            }
        }
```

`name` - the name of the supplier
`source` - url from where the data can be fetched
`response_format` - response_format is to map the supplier source params to our desired params

We would be able to map the supplier response params to the params desired for our data merging.
At the end of the procurement phase, we would have multiple supplier data following the same format.


### Supplier Data Merging Strategy

Here, we would have data in a common format across multiple suppliers.
I have used the following logic to merge the data

1. id - no logic
2. destination_id - no logic
3. name - the name with the longest length among the suppliers was chosen
4. description - the description with the longest length among the suppliers was chosen
5. lat - any supplier with a latitude param was chosen, if none had it the result will be 0
6. lng - any supplier with a longitude param was chosen, if none had it the result will be 0
7. address - the address param with the longest length among the suppliers was chosen
8. city - the city param with the longest length among the suppliers was chosen
9. country - the country param with the longest length among the suppliers was chosen
10. amenities - this was split into two. general and rooms. we have defined a map lookup in the config.json for general and room amenities. The data procured is looked up in the config and segregated to the corresponding bucket.
11. images - all three buckets (site, amenities, room) were procured , merged and presented
12. boooking_conditions - the boooking_conditions param with the longest length among the suppliers was chosen
