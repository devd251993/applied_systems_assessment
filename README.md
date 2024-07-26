# applied_systems_assessment

## Running the code base

Run following command <br></br>
`go build -o assessment ./cmd && ./assessment` <br></br>

This will start the server on port `:8080` <br> </br>

The server can serve following API's <br>

1. CreateGraph
2. AddEdgeToNodes
3. GetGraphDetails
4. GetShortesPath
5. DeleteGraph

<br>
<h2>
    1. CreateGraph
</h2>

This API creates a diconnected graph with specified number of verties.

```
- Method: POST
- Endpoint: /graph/create
- Query:
    - verties: int | number of nodes of graph. Should be a non zero number.
- Response:
    - success:
        {
            "Graph created with id: 1"
        }
```

<br> <br>

## 2. AddEdgeToNodes

This API adds a node to the existing graph.

```
- Method: PUT
- Endpoint: /graph/add-edge
- Query:
    - graphId: int (required) | id of graph.
    - source: int | start node for edge
    - destination: int | end node for edge
- Responses:
    - success:
        - return 200
    - Failure:
        - returns status code 404 if graph is not present
```

<br><br>

## 3. GetGraphDetails

This API returns details of the graph.

```
- Method: GET
- Endpoint: /graph/get
- Query:
    - graphId: int (required) | id of graph.
- Responses:
    - success:
        -   {
                "Reprentation" : "3 [{0:1}:1 {1:2}:1]",
                "NoOfNodes": 3
            }
            Note: The representation gives us number of nodes in the graph followed by edges present in the graph.
    - Failure:
        - returns status code 404 if graph is not present
```

<br><br>

## 4. GetShortesPath

This API returns shortest path between two nodes of graph.

```
- Method: GET
- Endpoint: /graph/shortest-path
- Query:
    - graphId: int (required) | id of graph.
    - source: int | start node for edge
    - destination: int | end node for edge
- Responses:
    - success:
        -   {
                Path : [0, 1, 2],
                Distance: 2
            }
    - Failure:
        - returns status code 404 if graph is not present
```

<br><br>

## 5. DeleteGraph

This API deletes the existing graph.

```
- Method: DELETE
- Endpoint: /graph/delete
- Query:
    - graphId: int (required) | id of graph.
- Responses:
    - success:
        - return 200
    - Failure:
        - returns status code 404 if graph is not present
```

<br><br>

