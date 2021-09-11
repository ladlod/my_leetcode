package common_list;

public class leetcode4 {
    public static void main(String args[]) {
        int[] nums1 = { 3, 4 };
        int[] nums2 = { 1, 2 };
        double ret = findMedianSortedArrays(nums1, nums2);
        System.out.println(ret);
    }

    public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int i = 0, j = 0, k = 0, h = 0;
        int tag = 1;
        double ret = 0;
        if ((nums1.length + nums2.length) % 2 == 0)
            tag = 0;
        while (i < nums1.length || j < nums2.length) {
            if (j >= nums2.length || (i < nums1.length && nums1[i] <= nums2[j])) {
                if (k == (nums1.length + nums2.length) / 2) {
                    if (tag == 1)
                        ret = nums1[i];
                    else
                        ret = (double) (nums1[i] + h) / 2;
                    break;
                }
                h = nums1[i];
                i++;
                k++;
            } else {
                if (k == (nums1.length + nums2.length) / 2) {
                    if (tag == 1)
                        ret = nums2[j];
                    else
                        ret = (double) (nums2[j] + h) / 2;
                    break;
                }
                h = nums2[j];
                j++;
                k++;
            }
        }
        /*
         * System.out.println("i=" + i); System.out.println("j=" + j);
         * System.out.println("k=" + k); System.out.println("h=" + h);
         */
        return ret;
    }
}