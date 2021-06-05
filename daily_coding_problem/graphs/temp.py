from collections import defaultdict

graph = {
    1 : [2,3],
    2 : [],
    3 : [4,5],
    4 : [6,7,8],
    5 :[],
    6 : [],
    7 : [],
    8 : []
}


def traverse(graph, curr, result):
    descendants = 0
    print(result)
    for child in graph[curr]:
        print(child)
        num_nodes, result = traverse(graph , child, result)
        print("num_node",num_nodes,result)
        result[child] += num_nodes - 1
        descendants += num_nodes

    return descendants+1, result

def max_edges(graph):
    start = list(graph)[0]
    vertices = defaultdict(int)

    _,descendants = traverse(graph, start, vertices)
    print(descendants.values())
    return len([val for val in descendants.values() if val%2==1])


print(max_edges(graph))
