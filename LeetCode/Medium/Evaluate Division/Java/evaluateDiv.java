class Solution {
    public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
        Map<String, Map<String, Double>> g = new HashMap<>();
        for (int i=0; i<equations.size(); i++) {
            List<String> equation = equations.get(i);
            String dividend = equation.get(0);
            String divisor = equation.get(1);
            double quotient = values[i];
            g.computeIfAbsent(dividend, k->new HashMap<>()).put(divisor, quotient);
            g.computeIfAbsent(divisor, k->new HashMap<>()).put(dividend, 1/quotient);
        }

        double[] results = new double[queries.size()];
        for (int i=0; i<queries.size(); i++) {
            List<String> query = queries.get(i);
            String dividend = query.get(0), divisor = query.get(1);
            if (!g.containsKey(dividend) || !g.containsKey(divisor)) results[i] = -1.0;
            else if (dividend == divisor) results[i] = 1.0;
            else {
                Set<String> visited = new HashSet<>();
                results[i] = calculate(g, dividend, divisor, visited, 1);
            }
        }
        return results;
    }

    private double calculate(Map<String, Map<String, Double>> g, String currNode, String targetNode, Set<String> visited, double prod) {
        visited.add(currNode);
        double ret = -1.0;
        Map<String, Double> nei = g.get(currNode);
        if (nei.containsKey(targetNode)) ret = prod * nei.get(targetNode);
        else {
            for (Map.Entry<String, Double> pair : nei.entrySet()) {
                String nextNode = pair.getKey();
                if (visited.contains(nextNode)) continue;
                ret = calculate(g, nextNode, targetNode, visited, prod*pair.getValue());
                if (ret != -1.0) break;
            }
        }
        visited.remove(currNode); // for next backtracking
        return ret;
    }

}
