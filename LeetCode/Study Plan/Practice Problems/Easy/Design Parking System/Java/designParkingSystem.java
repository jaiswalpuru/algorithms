class ParkingSystem {
    private final int big = 1, med = 2, small = 3;
    int b, m, s;
    public ParkingSystem(int big, int medium, int small) {
        b = big;
        m = medium;
        s = small;
    }
    
    public boolean addCar(int carType) {
        if (carType == big && b > 0) {
            b--;
            return true;
        }
        if (carType == small && s > 0) {
            s--;
            return true;
        }
        if (carType == med && m > 0) {
            m--;
            return true;
        }
        return false;
    }
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * ParkingSystem obj = new ParkingSystem(big, medium, small);
 * boolean param_1 = obj.addCar(carType);
 */
