from ctypes import *
from numpy.ctypeslib import ndpointer
import time
import matplotlib.pyplot as plt
import numpy as np
# plt.rcParams['axe.unicode_minus'] = False   #able to show the negative

lib = cdll.LoadLibrary('./PcieLaneMarginCmds.so')
lib2 = cdll.LoadLibrary('./LaneMargin.so')


def getRange(array, limit):
    count = 0
    for num in array:
        if num > limit:
            break
        count += 1
    return count


if __name__ == "__main__":
    # GetLaneMargin.argtypes =
    # bus = c_uint8(0xa)
    # device = c_uint8(0x0)
    # function = c_uint8(0x0)
    bus = 0xa
    device = 0
    function = 0
    lanenum = 0x1
    lib.GetVoltageSupported.argtypes = [c_int8,
                                        c_int8,
                                        c_int8,
                                        c_int,
                                        c_int]
    lib.GetVoltageSupported.restype = c_bool
    VoltageSupported = lib.GetVoltageSupported(
        bus, device, function, lanenum, 0x6)

    lib.ReportNumTimingSteps.argtypes = [c_uint8,
                                         c_uint8,
                                         c_uint8,
                                         c_uint,
                                         c_uint]
    lib.ReportNumTimingSteps.restype = c_uint8
    result = lib.ReportNumTimingSteps(bus, device, function, lanenum, 0x6)
    print('result = 0x%x' % result)
    if VoltageSupported:
        lib.ReportNumVoltageSteps.argtypes = [c_uint8,
                                              c_uint8,
                                              c_uint8,
                                              c_uint,
                                              c_uint]
        lib.ReportNumVoltageSteps.restype = c_uint8
        result2 = lib.ReportNumVoltageSteps(
            bus, device, function, lanenum, 0x6)
    else:
        result2 = 6
    print('result2 = 0x%x' % result2)
    ###############################################################
    start = time.time()
    for i in range(10):
        lib2.GetPcieLaneMarginingLeftTimingOffset.argtypes = [
            c_uint8, c_uint8, c_uint8, c_uint, c_uint, c_uint]
        lib2.GetPcieLaneMarginingLeftTimingOffset.restype = ndpointer(
            dtype=c_uint8, shape=(result,))
        LEFT = lib2.GetPcieLaneMarginingLeftTimingOffset(
            bus, device, function, lanenum, 0x6, 0x10)

        lib2.GetPcieLaneMarginingRightTimingOffset.argtypes = [
            c_uint8, c_uint8, c_uint8, c_uint, c_uint, c_uint]
        lib2.GetPcieLaneMarginingRightTimingOffset.restype = ndpointer(
            dtype=c_uint8, shape=(result,))
        RIGHT = lib2.GetPcieLaneMarginingRightTimingOffset(
            bus, device, function, lanenum, 0x6, 0x10)

        lib2.GetPcieLaneMarginingUpVoltageOffset.argtypes = [
            c_uint8, c_uint8, c_uint8, c_uint, c_uint, c_uint]
        lib2.GetPcieLaneMarginingUpVoltageOffset.restype = ndpointer(
            dtype=c_uint8, shape=(result2,))
        UP = lib2.GetPcieLaneMarginingUpVoltageOffset(
            bus, device, function, lanenum, 0x6, 0x10)

        lib2.GetPcieLaneMarginingDownVoltageOffset.argtypes = [
            c_uint8, c_uint8, c_uint8, c_uint, c_uint, c_uint]
        lib2.GetPcieLaneMarginingDownVoltageOffset.restype = ndpointer(
            dtype=c_uint8, shape=(result2,))
        DOWN = lib2.GetPcieLaneMarginingDownVoltageOffset(
            bus, device, function, lanenum, 0x6, 0x10)
        print(LEFT)
        print(RIGHT)
        print(UP)
        print(DOWN)

        leftNum = 0-getRange(LEFT, 2)
        rightNum = getRange(RIGHT, 2)
        upNum = 0-getRange(UP, 5)
        downNum = getRange(DOWN, 5)
        print(leftNum, rightNum, upNum, downNum)
        # X = np.linepave(-10,10,256,endpoint=True)
        # cos = np.cos(X)
        # points=[[leftNum,0],[0,upNum],[rightNum,0],[0,downNum]]
        # plt.plot([leftNum,upNum],[upNum,rightNum])
        # plt.plot([upNum,rightNum],[rightNum,downNum])
        # plt.plot([rightNum,downNum],[downNum,leftNum])
        # plt.plot([downNum,leftNum],[leftNum,upNum])
        # plt.plot([leftNum,0],[0,upNum],[rightNum,0],[0,downNum])
        x = [leftNum, 0, rightNum, 0]
        y = [0, upNum, 0, downNum]
        x.append(x[0])
        y.append(y[0])
        # plt.fill(x,y,facecolor='g',alpha=0.5)

        # plt.plot(x, y, color='red', lw=0.8)
        plt.plot(x, y, lw=0.8)

    end = time.time()
    print("it takes %.2f sec" % (end-start))
    plt.show()
